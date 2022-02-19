{#if connectionStore.validate(connection)}
    <Grid>
        <SubNav
                bind:mappings={mappings}
                bind:profileSelected="{profileSelected}"
                bind:pluginLink="{pluginLink}"
                bind:pluginField="{pluginField}"
                bind:template="{template}"
        />
        {#if profileSelected}
            <Camera
                    success={success}
                    bind:files="{files}"
            />
            {#if files}
                <Form
                        mappings={mappings}
                        pluginLink="{pluginLink}"
                        pluginField="{pluginField}"
                        template="{template}"
                        files="{files}"
                        bind:success="{success}"
                />
            {/if}
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
    let pluginLink = "";
    let pluginField = "";
    let template = "";
    let files = "";
    let success = "";

    let mappings = [];

    let connection = connectionStore.init();
    onMount(async () => connection = await connectionStore.mount());
</script>