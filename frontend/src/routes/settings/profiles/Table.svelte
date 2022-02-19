<Row>
    <Column style="padding-bottom: 10px">
        <StructuredList>
            <StructuredListHead>
                <StructuredListRow head>
                    <StructuredListCell head>Name</StructuredListCell>
                    <StructuredListCell head>Plugin</StructuredListCell>
                    <StructuredListCell head>Mappings</StructuredListCell>
                    <StructuredListCell head>Template</StructuredListCell>
                    <StructuredListCell head>Remove</StructuredListCell>
                </StructuredListRow>
            </StructuredListHead>
            <StructuredListBody>
                {#each profiles as profile}
                    <TableRow
                            profile="{profile}"
                            handleDelete="{handleDelete}"
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
    import TableRow from "./TableRow.svelte";
    import {profileListStore} from "$lib/stores/profiles";
    import {errorStore} from "$lib/stores/error";

    export let profiles;

    /** handleDelete delete profile from list */
    const handleDelete = (profile) => {
        if (profiles.length === 1) {
            errorStore.set({
                title: "Cannot remove",
                error: "Cannot delete if single profile remaining",
            });
            return;
        }
        profiles = profiles.filter((p) => {
            return p.name !== profile.name;
        });
        profileListStore.persist(profiles);
    }

    /** handleUpdate Update and set persistence */
    const handleUpdate = () => {
        profileListStore.persist(profiles);
    }
</script>