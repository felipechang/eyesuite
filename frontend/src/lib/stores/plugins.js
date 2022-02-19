import { writable } from "svelte/store";
const init = [];
const ENDPOINT = "http://localhost:5000/plugins";
function createStore() {
    const { subscribe, set } = writable(init);
    return {
        findOne: (plugins, name) => {
            const p = plugins.find((e) => {
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
                };
            }
            return p;
        },
        init: () => init,
        mount: async () => {
            const response = await fetch(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (plugins) => {
            const response = await fetch(ENDPOINT, {
                method: 'POST',
                body: JSON.stringify(plugins)
            });
            const data = await response.json();
            set(data.data);
        },
    };
}
export const pluginListStore = createStore();
//# sourceMappingURL=plugins.js.map