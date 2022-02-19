import {writable} from "svelte/store";

const init: IPluginInit[] = [];

const ENDPOINT = "http://localhost:5000/plugins";

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
            const response = await fetch(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (plugins: IPluginInit[]) => {
            const response = await fetch(ENDPOINT, {
                method: 'POST',
                body: JSON.stringify(plugins)
            })
            const data = await response.json();
            set(data.data);
        },
    }
}

export const pluginListStore = createStore();