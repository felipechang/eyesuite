import { writable } from "svelte/store";
import { getServer, postServer } from "$lib/stores/wrapper";
const init = [];
const ENDPOINT = "/profiles";
function createStore() {
    const { subscribe, set } = writable(init);
    return {
        findOne: (profiles, name) => {
            const p = profiles.find((e) => {
                return e.name === name;
            });
            if (!p) {
                return { name: "", plugin: "", template: "", mapping: [] };
            }
            return p;
        },
        init: () => init,
        mount: async () => {
            const response = await getServer(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (profiles) => {
            const response = await postServer(ENDPOINT, JSON.stringify(profiles));
            const data = await response.json();
            set(data.data);
        },
    };
}
export const profileListStore = createStore();
//# sourceMappingURL=profiles.js.map