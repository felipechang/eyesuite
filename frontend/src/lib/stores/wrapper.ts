import {loginStore} from "$lib/stores/login";

const postServer = async (url: string, body: any): Promise<Response> => {
    const response = await fetch(`${url}`, {
        headers: makeHeaders(),
        method: 'POST',
        body: body
    });
    if (response.status === 401) {
        await loginStore.refresh()
        return postServer(url, body);
    }
    return response;
}

const getServer = async (url: string): Promise<Response> => {
    const response = await fetch(`${url}`, {
        headers: makeHeaders(),
    });
    if (response.status === 401) {
        await loginStore.refresh()
        return getServer(url);
    }
    return response;
}

const loginServer = async (username: string, password: string): Promise<Response> => {
    return await fetch(`api/login`, {
        headers: {
            "Authorization": "",
        },
        method: 'POST',
        body: JSON.stringify({
            username: username,
            password: password,
        })
    });
}

const makeHeaders = () => {
    const auth = localStorage.getItem("auth") as string;
    const tokens = auth ? JSON.parse(auth) : {"access_token": "", "refresh_token": ""};
    return {
        "Authorization": `${tokens.refresh_token} ${tokens.access_token}`,
    }
}

export {
    loginServer,
    postServer,
    getServer,
}