<Row>
    <Column style="padding-bottom: 10px">
        <StructuredList>
            <StructuredListHead>
                <StructuredListRow head>
                    <StructuredListCell head>Username</StructuredListCell>
                    <StructuredListCell head>Full Name</StructuredListCell>
                    <StructuredListCell head>Enabled</StructuredListCell>
                    <StructuredListCell head>Administrator</StructuredListCell>
                    <StructuredListCell head>Key</StructuredListCell>
                    <StructuredListCell head>Remove</StructuredListCell>
                </StructuredListRow>
            </StructuredListHead>
            <StructuredListBody>
                {#each users as user}
                    <TableRow
                            users="{users}"
                            user="{user}"
                            handleDelete="{handleDelete}"
                    />
                {/each}
            </StructuredListBody>
        </StructuredList>
    </Column>
</Row>

<script lang="ts">
    import {
        Column,
        Row,
        StructuredList,
        StructuredListBody,
        StructuredListCell,
        StructuredListHead,
        StructuredListRow,
    } from "carbon-components-svelte";
    import {errorStore} from "$lib/stores/error";
    import {usersStore} from "$lib/stores/users";
    import TableRow from "./TableRow.svelte";

    export let users;

    /** handleDelete delete user from list */
    const handleDelete = (user) => {
        if (users.length === 1) {
            errorStore.set({
                title: "Cannot remove",
                error: "Cannot delete if single user remaining",
            });
            return;
        }
        users = users.filter((u) => {
            return u.username !== user.username;
        });
        usersStore.persist(users);
    }
</script>

<style>
    :global(svg.clickable) {
        cursor: pointer;
    }
</style>