<!--Camera input-->
<Row>
    <Column/>
    <Column>
        {#if success !== ""}
            <InlineNotification
                    kind="success"
                    title="Success:"
                    subtitle={success}
                    on:close={handleMessageClose}
            />
        {/if}
        <div on:click={fileClick} style="text-align: center">
            {#if avatar}
                <img style="max-width: 500px" src="{avatar}" alt="Avatar"/>
            {:else}
                <CameraPic style="cursor: pointer"/>
            {/if}
        </div>
    </Column>
    <Column/>
</Row>

<!--Actual input-->
<input accept=".jpg, .jpeg, .png" bind:this={fileInput}
       name="file"
       on:change={onImageChangeHandler}
       style="display: none"
       type="file"/>

<script>
    import CameraPic from "carbon-pictograms-svelte/lib/Camera.svelte";
    import {Column, InlineNotification, Row} from "carbon-components-svelte";

    export let files;
    export let success;

    let avatar;
    $:avatar = success ? "" : avatar;

    let fileInput;

    const handleMessageClose = () => {
        success = "";
    }

    /** fileClick Simulate a file click */
    const fileClick = () => {
        fileInput.click();
    }

    /** onImageChangeHandler Make sure we have a file and set avatar */
    const onImageChangeHandler = (e) => {
        files = e.target.files[0];
        if (!files) {
            return;
        }
        avatar = "";
        let reader = new FileReader();
        reader.readAsDataURL(files);
        reader.onload = (e) => {
            avatar = e.target.result;
        };
    }

</script>