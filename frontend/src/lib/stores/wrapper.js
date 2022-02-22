import { loginStore } from "$lib/stores/login";
const postServer = async (url, body) => {
    const headers = makeHeaders();
    headers["Content-Type"] = "application/json";
    const response = await fetch(`${url}`, {
        headers,
        method: 'POST',
        body: body
    });
    if (response.status === 401) {
        localStorage.removeItem("auth");
        await loginStore.refresh();
        return postServer(url, body);
    }
    return response;
};
const getServer = async (url) => {
    const headers = makeHeaders();
    headers["Content-Type"] = "application/json";
    const response = await fetch(`${url}`, { headers });
    if (response.status === 401) {
        localStorage.removeItem("auth");
        await loginStore.refresh();
        return getServer(url);
    }
    return response;
};
const makeHeaders = () => {
    const auth = localStorage.getItem("auth");
    const tokens = auth ? JSON.parse(auth) : { "access_token": "", "refresh_token": "" };
    return {
        "Authorization": `${tokens.refresh_token} ${tokens.access_token}`,
    };
};
export { makeHeaders, postServer, getServer, };
//# sourceMappingURL=wrapper.js.map