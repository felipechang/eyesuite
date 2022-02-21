import {loginStore} from "$lib/stores/login";

const postServer = async (url: string, body: BodyInit): Promise<Response> => {
    const response = await fetch(`${url}`, {
        headers: makeHeaders(),
        method: 'POST',
        body: body
    });
    if (response.status === 401) {
        localStorage.removeItem("auth");
        await loginStore.refresh();
        return postServer(url, body);
    }
    return response;
}

const getServer = async (url: string): Promise<Response> => {
    const response = await fetch(`${url}`, {
        headers: makeHeaders(),
    });
    if (response.status === 401) {
        localStorage.removeItem("auth");
        await loginStore.refresh()
        return getServer(url);
    }
    return response;
}

const makeHeaders = () => {
    const auth = localStorage.getItem("auth") as string;
    const tokens = auth ? JSON.parse(auth) : {"access_token": "", "refresh_token": ""};
    return {
        "Authorization": `${tokens.refresh_token} ${tokens.access_token}`,
    }
}

export {
    makeHeaders,
    postServer,
    getServer,
}