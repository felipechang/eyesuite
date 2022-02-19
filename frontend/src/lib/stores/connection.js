import { writable } from "svelte/store";
const init = {
    version: "",
    ws_url: "",
    realm: "",
    consumer_key: "",
    consumer_secret: "",
    token_id: "",
    token_secret: "",
};
const ENDPOINT = "http://localhost:5000/config";
function createStore() {
    const { subscribe, set } = writable(init);
    return {
        init: () => init,
        mount: async () => {
            const response = await fetch(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (connection) => {
            const response = await fetch(ENDPOINT, {
                method: 'POST',
                body: JSON.stringify(connection)
            });
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