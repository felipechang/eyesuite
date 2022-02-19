import { writable } from "svelte/store";
const init = [];
const ENDPOINT = "http://localhost:5000/profiles";
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
            const response = await fetch(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (profiles) => {
            const response = await fetch(ENDPOINT, {
                method: 'POST',
                body: JSON.stringify(profiles)
            });
            const data = await response.json();
            set(data.data);
        },
    };
}
export const profileListStore = createStore();
//# sourceMappingURL=profiles.js.map