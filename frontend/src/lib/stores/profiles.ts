import {writable} from "svelte/store";

const init: IProfileInit[] = [];

const ENDPOINT = "http://localhost:5000/profiles";

function createStore() {
    const {subscribe, set} = writable(init);
    return {
        findOne: (profiles: IProfileInit[], name: string): IProfileInit => {
            const p = profiles.find((e: IProfileInit) => {
                return e.name === name;
            });
            if (!p) {
                return {name: "", plugin: "", template: "", mapping: []}
            }
            return p;
        },
        init: (): IProfileInit[] => init,
        mount: async () => {
            const response = await fetch(ENDPOINT);
            const data = await response.json();
            return data.data;
        },
        subscribe,
        persist: async (profiles: IProfileInit[]) => {
            const response = await fetch(ENDPOINT, {
                method: 'POST',
                body: JSON.stringify(profiles)
            })
            const data = await response.json();
            set(data.data);
        },
    };
}

export const profileListStore = createStore();