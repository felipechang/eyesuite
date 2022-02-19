<Grid>
    <Title
            bind:open={open}
            details="Manage Reader profile options"
            includeIcon
            name="Reader Profiles"
    />
    <Table profiles="{profiles}"/>
    <ModalForm profiles="{profiles}" bind:open={open}/>
</Grid>

<script lang="ts">
    import {Grid,} from "carbon-components-svelte";
    import Table from "./Table.svelte";
    import Title from "$lib/components/Title.svelte";
    import ModalForm from "./ModalForm.svelte";
    import {profileListStore} from "$lib/stores/profiles";
    import {onMount} from "svelte";

    let open = false;

    let profiles = profileListStore.init();
    profileListStore.subscribe((p) => profiles = p);
    onMount(async () => profiles = await profileListStore.mount());
</script>