import {writable} from "svelte/store";
import {getServer, postServer} from "$lib/stores/wrapper";

const init: IPluginInit[] = [];

const ENDPOINT = "/api/plugins";

function createStore() {
    const {subscribe, set} = writable(init);
    return {
        findOne: (plugins: IPluginInit[], name: string): IPluginInit => {
            const p = plugins.find((e: IPluginInit) => {
                return e.name === name;
            });
            if (!p) {
                return {
                    name: "",
                    description: "",
                    endpoint: "",
                    endpointField: "",
                    enabled: false,
                    mapping: []
                }
            }
            return p;
        },
        init: (): IPluginInit[] => init,
        mount: async () => {
            const response = await getServer(ENDPOINT);
            return await response.json();
        },
        subscribe,
        persist: async (plugins: IPluginInit[]) => {
            const response = await postServer(ENDPOINT, JSON.stringify(plugins));
            set(await response.json());
        },
    }
}

export const pluginListStore = createStore();