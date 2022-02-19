<!--Camera input-->
<Row>
    <Column/>
    <Column>
        {#if avatar}
            <ImageLoader fadeIn src="{avatar}"/>
        {:else}
            <div on:click={fileClick} style="text-align: center">
                <CameraPic style="cursor: pointer"/>
            </div>
        {/if}
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
    import {Column, ImageLoader, Row} from "carbon-components-svelte";

    let avatar;
    let files;
    let fileInput;

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
        let reader = new FileReader();
        reader.readAsDataURL(files);
        reader.onload = (e) => {
            avatar = e.target.result;
        };
    }

</script>