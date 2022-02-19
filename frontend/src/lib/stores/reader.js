function createStore() {
    return {
        submit: async (template, files) => {
            const formData = new FormData();
            formData.append("file", files);
            formData.append("template", template);
            const response = await fetch(`http://localhost:5000/postImage`, {
                method: 'POST',
                body: formData
            });
            // invalid credentials
            if (response.status === 400) {
                return {
                    error: true,
                    msg: "Invalid NetSuite credentials. Please contact your administrator."
                };
            }
            const res = await response.json();
            return { error: false, msg: res.data };
        },
        read: async (link, files) => {
            const formData = new FormData();
            formData.append("file", files);
            const response = await fetch(`http://localhost:5000${link}`, {
                method: 'POST',
                body: formData
            });
            const data = await response.json();
            return data.data;
        }
    };
}
export const readerStore = createStore();
//# sourceMappingURL=reader.js.map