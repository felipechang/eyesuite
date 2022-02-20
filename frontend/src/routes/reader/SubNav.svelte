<Row>
    <Column>
        <Title
                details="Take a photo and fill in the details"
                name="Reader"
        />
    </Column>
    <Column>
        <Select
                bind:selected={current}
                labelText="Select Profile:"
                on:input={handleSelectChange}
        >
            <SelectItem text="" value=""/>
            {#each profiles as profile}
                <SelectItem text="{profile.name}" value="{profile.name}"/>
            {/each}
        </Select>
    </Column>
</Row>

<script lang="ts">
    import {Column, Row, Select, SelectItem} from "carbon-components-svelte";
    import Title from "$lib/components/Title.svelte";
    import {onMount} from "svelte";

    export let profiles;
    export let loadMappings;
    export let updateSelectChange;

    let current;

    onMount(async () => {
        await loadMappings();
        current = localStorage.getItem("profile-selected");
        updateSelectChange(current);
    });

    /** handleSelectChange export selected profile name */
    const handleSelectChange = (e) => {
        localStorage.setItem("profile-selected", e.target.value);
        updateSelectChange(e.target.value);
    }
</script>