<Content>
    <Form>
        <Grid>
            <Row>
                <Column/>
                <Column>
                    <h1>Welcome to EyeSuite</h1>
                    <div style="padding: 10px 0 10px 0">
                        <p>Enter login details:</p>
                    </div>
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
    import {loginStore} from "$lib/stores/login.js";

    export let handleLogIn;
    export let isAdministrator;

    let username = "";
    let password = "";

    const handleServerLogin = async () => {
        const res = await loginStore.login(username, password);
        isAdministrator = btoa(username + "1") === res.control;
        localStorage.setItem("auth", JSON.stringify(res));
        localStorage.setItem("user", JSON.stringify({username: username, lastLogin: (new Date()).getTime()}));
        handleLogIn();
    }
</script>