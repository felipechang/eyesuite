import {writable} from "svelte/store";

const init: IUserInit[] = [];

const ENDPOINT = "http://localhost:5000/users";

function createStore() {
    const {subscribe, set} = writable(init);
    return {
        init: (): IUserInit[] => init,
        mount: async () => {
            const response = await fetch(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (users: IUserInit[]) => {
            const response = await fetch(ENDPOINT, {
                method: 'POST',
                body: JSON.stringify(users)
            })
            const data = await response.json();
            set(data.data);
        },
    };
}

export const usersStore = createStore();