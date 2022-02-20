<Grid>
    <Title
            bind:open={open}
            details="Manage user access"
            includeIcon
            name="Users"
    />
    <Table users="{users}"/>
    <ModalForm bind:open={open} users="{users}"/>
</Grid>

<script lang="ts">
    import {Grid,} from "carbon-components-svelte";
    import Table from "./Table.svelte";
    import Title from "$lib/components/Title.svelte";
    import ModalForm from "./ModalForm.svelte";
    import {usersStore} from "$lib/stores/users";
    import {onMount} from "svelte";

    let open = false;

    let users = usersStore.init();
    usersStore.subscribe((e) => users = e);
    onMount(async () => users = await usersStore.mount());
</script>