<Modal
        bind:open
        modalHeading="Add User"
        on:click:button--secondary={()=>(open=false)}
        on:close={handleClose}
        on:open
        on:submit={handleSubmit}
        primaryButtonText="Add"
        secondaryButtonText="Cancel"
>

    <FluidForm>
        <Grid>
            <Row>
                <Column style="padding-bottom: 10px">
                    <TextInput
                            labelText="Username"
                            bind:value={user.username}
                            on:change={handleUsernameChange}
                            required
                    />
                </Column>
            </Row>

            <Row>
                <Column style="padding-bottom: 10px">
                    <TextInput
                            labelText="Full Name"
                            bind:value={user.name}
                            required
                    />
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    <PasswordInput
                            labelText="Password"
                            bind:value={user.password}
                            required
                            type="password"
                    />
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    <Checkbox
                            labelText="Enabled"
                            checked="{user.enabled ? 'checked': ''}"
                            on:change={handleEnabledChange}/>
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    <Checkbox
                            labelText="Administrator"
                            checked="{user.admin ? 'checked': ''}"
                            on:change={handleAdminChange}/>
                </Column>
            </Row>
        </Grid>
    </FluidForm>
</Modal>

<script lang="ts">
    import {Checkbox, Column, FluidForm, Grid, Modal, PasswordInput, Row, TextInput} from "carbon-components-svelte";
    import {errorStore} from "$lib/stores/error";
    import {usersStore} from "$lib/stores/users";

    export let open;

    export let users;

    let user = {
        name: "",
        username: "",
        key: "",
        password: "",
        enabled: false,
        admin: false,
    };

    /** handleClose Reset values and close modal */
    const handleClose = () => {
        user.name = "";
        user.username = "";
        user.key = "";
        user.password = "";
        user.enabled = false;
        user.admin = false;
    }

    /** handleSubmit Add user and close modal */
    const handleSubmit = () => {
        users.push({
            name: user.name,
            username: user.username,
            key: user.key,
            enabled: user.enabled,
            admin: user.admin,
            password: user.password,
        });
        open = false;
        usersStore.persist(users);
    }

    /** handleUsernameChange Show a duplicate error if username is repeated */
    const handleUsernameChange = () => {
        if (!users) {
            return;
        }
        user.username = user.username.trim();
        for (let i = 0; i < users.length; i++) {
            if (users[i].username.trim() === user.username) {
                errorStore.set({
                    title: "Duplicate record",
                    error: "Usernames must be unique",
                });
                user.username = "";
                return;
            }
        }
    }

    /** handleEnabledChange Flip enabled value */
    const handleEnabledChange = () => {
        user.enabled = !user.enabled;
    }

    /** handleAdminChange Flip admin value */
    const handleAdminChange = () => {
        user.admin = !user.admin;
    }

</script>