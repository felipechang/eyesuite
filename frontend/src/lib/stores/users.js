import { writable } from "svelte/store";
import { getServer, postServer } from "$lib/stores/wrapper";
const init = [];
const ENDPOINT = "/api/users";
function createStore() {
    const { subscribe, set } = writable(init);
    return {
        init: () => init,
        mount: async () => {
            const response = await getServer(ENDPOINT);
            return await response.json();
        },
        subscribe,
        persist: async (users) => {
            const response = await postServer(ENDPOINT, JSON.stringify(users));
            set(await response.json());
        },
    };
}
export const usersStore = createStore();
//# sourceMappingURL=users.js.map