import { writable } from "svelte/store";
import { getServer, postServer } from "$lib/stores/wrapper";
const init = {
    version: "",
    ws_url: "",
    realm: "",
    consumer_key: "",
    consumer_secret: "",
    token_id: "",
    token_secret: "",
};
const ENDPOINT = "/api/config";
function createStore() {
    const { subscribe, set } = writable(init);
    return {
        init: () => init,
        mount: async () => {
            console.log("connectionStore.mount");
            const response = await getServer(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (connection) => {
            const response = await postServer(ENDPOINT, JSON.stringify(connection));
            const data = await response.json();
            set(data.data);
        },
        validate: (connection) => {
            return (connection.version !== "" &&
                connection.ws_url !== "" &&
                connection.realm !== "" &&
                connection.consumer_key !== "" &&
                connection.consumer_secret !== "" &&
                connection.token_id !== "" &&
                connection.token_secret !== "");
        }
    };
}
export const connectionStore = createStore();
//# sourceMappingURL=connection.js.map