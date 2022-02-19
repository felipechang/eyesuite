{#if connectionStore.validate(connection)}
    <Grid>
        <SubNav bind:mappings={mappings} bind:profileSelected="{profileSelected}"/>
        {#if profileSelected}
            <Camera/>
            <Form mappings={mappings}/>
        {/if}
    </Grid>
{:else}
    <p>Please complete NetSuite connection settings</p>
{/if}

<script lang="ts">
    import {connectionStore} from "$lib/stores/connection";
    import {Grid} from "carbon-components-svelte";
    import Camera from "./reader/Camera.svelte";
    import SubNav from "./reader/SubNav.svelte";
    import Form from "./reader/Form.svelte";
    import {onMount} from "svelte";

    let profileSelected = "";
    let mappings = [];
    let connection = connectionStore.init();
    connectionStore.subscribe((c) => connection = c);
    onMount(async () => connection = await connectionStore.mount());

</script>