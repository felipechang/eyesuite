<Row>
    <Column style="padding-bottom: 10px">
        <StructuredList>
            <StructuredListHead>
                <StructuredListRow head>
                    <StructuredListCell head>Enable</StructuredListCell>
                    <StructuredListCell head>Name</StructuredListCell>
                    <StructuredListCell head>Description</StructuredListCell>
                    <StructuredListCell head>Mappings</StructuredListCell>
                </StructuredListRow>
            </StructuredListHead>
            <StructuredListBody>
                {#each plugins as plugin}
                    <TableRow
                            plugin="{plugin}"
                            handleUpdate="{handleUpdate}"
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
    import {pluginListStore} from "$lib/stores/plugins";
    import TableRow from "./TableRow.svelte";
    import {onMount} from "svelte";

    let plugins = pluginListStore.init();
    pluginListStore.subscribe((p) => plugins = p);
    onMount(async () => plugins = await pluginListStore.mount());

    const handleUpdate = () => {
        pluginListStore.persist(plugins);
    }
</script>