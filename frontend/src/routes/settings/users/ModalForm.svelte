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
                            bind:value={user.username}
                            labelText="Username"
                            on:change={handleUsernameChange}
                            required
                    />
                </Column>
            </Row>

            <Row>
                <Column style="padding-bottom: 10px">
                    <TextInput
                            bind:value={user.name}
                            labelText="Full Name"
                            required
                    />
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    <PasswordInput
                            bind:value={user.password}
                            labelText="Password"
                            required
                            type="password"
                    />
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    <Checkbox
                            checked="{user.enabled ? 'checked': ''}"
                            labelText="Enabled"
                            on:change={handleEnabledChange}/>
                </Column>
            </Row>
            <Row>
                <Column style="padding-bottom: 10px">
                    <Checkbox
                            checked="{user.admin ? 'checked': ''}"
                            labelText="Administrator"
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