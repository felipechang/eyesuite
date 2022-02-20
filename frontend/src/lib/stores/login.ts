import {loginServer, postServer} from "$lib/stores/wrapper";

const LOGOUT = "/logout";
const REFRESH = "/refresh";

function createStore() {
    return {
        login: async (username: string, password: string) => {
            const response = await loginServer(username, password)
            return await response.json();
        },
        logout: async () => {
            const response = await postServer(LOGOUT, "");
            return await response.json();
        },
        refresh: async () => {
            const response = await postServer(REFRESH, "");
            const data = await response.json();
            const auth = localStorage.getItem("auth") as string;
            const tokens = auth ? JSON.parse(auth) : {"control": "", "access_token": "", "refresh_token": ""};
            tokens.access_token = data.data.access_token;
            tokens.refresh_token = data.data.refresh_token;
            localStorage.setItem("auth", JSON.stringify(tokens));
        },
    }
}

export const loginStore = createStore();