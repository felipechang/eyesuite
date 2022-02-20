import { loginStore } from "$lib/stores/login";
const baseUrl = "http://localhost:5000";
const postServer = async (url, body) => {
    const response = await fetch(`${baseUrl}${url}`, {
        headers: makeHeaders(),
        method: 'POST',
        body: body
    });
    if (response.status === 401) {
        await loginStore.refresh();
        return postServer(url, body);
    }
    return response;
};
const getServer = async (url) => {
    const response = await fetch(`${baseUrl}${url}`, {
        headers: makeHeaders(),
    });
    if (response.status === 401) {
        await loginStore.refresh();
        return getServer(url);
    }
    return response;
};
const loginServer = async (username, password) => {
    return await fetch(`${baseUrl}/login`, {
        headers: {
            "Authorization": "",
        },
        method: 'POST',
        body: JSON.stringify({
            username: username,
            password: password,
        })
    });
};
const makeHeaders = () => {
    const auth = localStorage.getItem("auth");
    const tokens = auth ? JSON.parse(auth) : { "access_token": "", "refresh_token": "" };
    return {
        "Authorization": `${tokens.refresh_token} ${tokens.access_token}`,
    };
};
export { loginServer, postServer, getServer, };
//# sourceMappingURL=wrapper.js.map