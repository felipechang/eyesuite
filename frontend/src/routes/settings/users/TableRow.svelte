<StructuredListRow>
    <StructuredListCell>{user.username}</StructuredListCell>
    <StructuredListCell>
        <TextInput bind:value={user.name} on:change={()=>usersStore.persist(users)} required/>
    </StructuredListCell>
    <StructuredListCell>
        <Checkbox checked="{user.enabled ? 'checked': ''}"
                  on:change={()=>handleEnableChange(user)}/>
    </StructuredListCell>
    <StructuredListCell>
        <Checkbox checked="{user.admin ? 'checked': ''}" on:change={()=>handleAdminChange(user)}/>
    </StructuredListCell>
    <StructuredListCell>{user.key}</StructuredListCell>
    <StructuredListCell>
        <TrashCan24 class="clickable" on:click={()=>handleDelete(user)}/>
    </StructuredListCell>
</StructuredListRow>


<script lang="ts">
    import {Checkbox, StructuredListCell, StructuredListRow, TextInput,} from "carbon-components-svelte";
    import {errorStore} from "$lib/stores/error";
    import {usersStore} from "$lib/stores/users";
    import {TrashCan24} from "carbon-icons-svelte";

    export let user;
    export let users;
    export let handleDelete;

    /** handleEnableChange Persist change when user is enabled */
    const handleEnableChange = (user) => {
        user.enabled = !user.enabled;
        const found = users.find((e) => {
            return e.enabled;
        });
        if (!found) {
            errorStore.set({
                title: "Cannot disable",
                error: "At least one user must be enabled",
            });
            user.enabled = !user.enabled;
        }
        usersStore.persist(users);
    }

    /** handleAdminChange Persist change when user is admin */
    const handleAdminChange = (user) => {
        user.admin = !user.admin;
        const found = users.find((e) => {
            return e.admin;
        });
        if (!found) {
            errorStore.set({
                title: "Cannot disable",
                error: "At least one user must be administrator",
            });
            user.admin = !user.admin;
        }
        usersStore.persist(users);
    }

</script>

<style>
    :global(svg.clickable) {
        cursor: pointer;
    }
</style>