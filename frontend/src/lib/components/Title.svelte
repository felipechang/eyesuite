<Row>
    <Column style="padding-bottom: 20px">
        <div>
            <span class="text">
                {name}
            </span>
            {#if includeIcon}
                <span class="icon">
                    <AddFilled24 class="clickable" on:click={clickHandler}/>
                </span>

            {/if}
        </div>
        <p>{details}</p>
    </Column>
</Row>

<script lang="ts">
    import {Column, Row} from "carbon-components-svelte";
    import {AddFilled24} from "carbon-icons-svelte";

    export let name;
    export let details;
    export let includeIcon
    export let open;

    const clickHandler = () => {
        open = true;
    }

</script>

<style>
    :global(svg.clickable) {
        cursor: pointer;
    }

    .icon, .text {
        vertical-align: middle;
        display: inline-block;
    }

    .icon {
        padding-left: 10px;
    }

    .text {
        font-size: 2em;
        padding-bottom: 10px;
    }
</style>