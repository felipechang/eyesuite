import { makeHeaders, postServer } from "$lib/stores/wrapper";
function createStore() {
    return {
        login: async (username, password) => {
            const response = await fetch(`/api/login`, {
                headers: {
                    "Authorization": "",
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
            const auth = localStorage.getItem("auth");
            localStorage.removeItem("auth");
            // call for new token
            const response = await postServer("/api/refresh", "");
            const data = await response.json();
            const tokens = auth ? JSON.parse(auth) : { "control": "", "access_token": "", "refresh_token": "" };
            tokens.access_token = data.data.access_token;
            tokens.refresh_token = data.data.refresh_token;
            // set new token
            localStorage.setItem("auth", JSON.stringify(tokens));
        },
    };
}
export const loginStore = createStore();
//# sourceMappingURL=login.js.map