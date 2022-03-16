import {makeHeaders, postServer} from "$lib/stores/wrapper";

function createStore() {
    return {
        login: async (username: string, password: string) => {
            const response = await fetch(`/api/login`, {
                headers: {
                    "Authorization": "",
                    "Content-Type": "application/json"
                },
                method: 'POST',
                body: JSON.stringify({
                    username: username,
                    password: password,
                })
            });
            return await response.json();
        },
        logout: async () => {
            const response = await fetch(`/api/logout`, {
                headers: makeHeaders(),
                method: 'POST',
                body: ""
            });
            return await response.json();
        },
        refresh: async () => {

            // get current and empty in case request fails
            const auth = localStorage.getItem("auth") as string;
            localStorage.removeItem("auth");

            // call for new token
            const response = await postServer("/api/refresh", "");
            if (response.status === 401) {
                localStorage.removeItem("auth");
                throw new Error("could not refresh token")
            }
            const data = await response.json();
            const tokens = auth ? JSON.parse(auth) : {"control": "", "access_token": "", "refresh_token": ""};
            tokens.access_token = data.access_token;
            tokens.refresh_token = data.refresh_token;

            // set new token
            localStorage.setItem("auth", JSON.stringify(tokens));
        },
    }
}

export const loginStore = createStore();