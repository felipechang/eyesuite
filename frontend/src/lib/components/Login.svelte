<Content>
    <Form>
        <Grid>
            <Row>
                <Column/>
                <Column>
                    <Title
                            details="Enter login details:"
                            name="Company Name"
                    />
                </Column>
                <Column/>
            </Row>
            <Row>
                <Column/>
                <Column>
                    <TextInput
                            bind:value={username}
                            labelText="User name"
                            placeholder="Enter user name..."
                            required
                    />
                </Column>
                <Column/>
            </Row>
            <Row>
                <Column/>
                <Column>
                    <br/>
                    <PasswordInput
                            bind:value={password}
                            labelText="Password"
                            placeholder="Enter password..."
                            required
                            type="password"
                    />
                </Column>
                <Column/>
            </Row>
            <Row>
                <Column/>
                <Column>
                    <br/>
                    {#if username && password}
                        <Button on:click={handleServerLogin}>Login</Button>
                    {/if}
                </Column>
                <Column/>
            </Row>
        </Grid>
    </Form>
</Content>

<script>
    import {Button, Column, Content, Form, Grid, PasswordInput, Row, TextInput} from "carbon-components-svelte";
    import Title from "$lib/components/Title.svelte";
    import {loginStore} from "$lib/stores/login.js";
    import {errorStore} from "$lib/stores/error.js";

    export let handleLogIn;
    export let isAdministrator;

    let username = "";
    let password = "";

    const handleServerLogin = async () => {
        const res = await loginStore.login(username, password);
        if (res.error) {
            errorStore.set({title: "Login Error", error: res.data.toUpperCase()})
        } else {
            isAdministrator = btoa(username + "1") === res.data.control;
            localStorage.setItem("auth", JSON.stringify(res.data));
            handleLogIn();
        }
    }
</script>