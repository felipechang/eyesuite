function createStore() {
    return {
        submit: (init: IReaderInit) => {
            console.log(init);
        }
    };
}

export const readerStore = createStore();