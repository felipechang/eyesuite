import {writable} from "svelte/store";
import {getServer, postServer} from "$lib/stores/wrapper";

const init: IConnectionInit = {
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
    const {subscribe, set} = writable(init);
    return {
        init: (): IConnectionInit => init,
        mount: async () => {
            const response = await getServer(ENDPOINT);
            return await response.json();
        },
        subscribe,
        persist: async (connection: IConnectionInit) => {
            const response = await postServer(ENDPOINT, JSON.stringify(connection));
            set(await response.json());
        },
        validate: (connection: IConnectionInit): boolean => {
            return (
                connection.version !== "" &&
                connection.ws_url !== "" &&
                connection.realm !== "" &&
                connection.consumer_key !== "" &&
                connection.consumer_secret !== "" &&
                connection.token_id !== "" &&
                connection.token_secret !== ""
            );
        }
    };
}

export const connectionStore = createStore();