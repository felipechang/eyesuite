<Row>
    <Column style="padding-bottom: 10px">
        <FluidForm>
            <Grid>
                <Row>
                    <Column style="padding-bottom: 10px">
                        <TextInput
                                bind:value={connection.version}
                                labelText="SuiteTalk Version"
                                on:change={()=>connectionStore.persist(connection)}
                                placeholder="2021_2"
                                required
                        />
                    </Column>
                </Row>
                <Row>
                    <Column style="padding-bottom: 10px">
                        <TextInput
                                bind:value={connection.ws_url}
                                labelText="Web Services URL"
                                on:change={()=>{
                                    handleLinkChange();
                                    connectionStore.persist(connection)
                                }}
                                placeholder="https://1234567-sb1.suitetalk.api.netsuite.com/"
                                required
                                type="url"
                        />
                    </Column>
                </Row>
                <Row>
                    <Column style="padding-bottom: 10px">
                        <PasswordInput
                                bind:value={connection.consumer_key}
                                labelText="Consumer Key"
                                on:change={()=>connectionStore.persist(connection)}
                                required
                                type="password"
                        />
                    </Column>
                </Row>
                <Row>
                    <Column style="padding-bottom: 10px">
                        <PasswordInput
                                bind:value={connection.consumer_secret}
                                labelText="Consumer Secret"
                                on:change={()=>connectionStore.persist(connection)}
                                required
                                type="password"
                        />
                    </Column>
                </Row>
                <Row>
                    <Column style="padding-bottom: 10px">
                        <PasswordInput
                                bind:value={connection.token_id}
                                labelText="Token ID"
                                on:change={()=>connectionStore.persist(connection)}
                                required
                                type="password"
                        />

                    </Column>
                </Row>
                <Row>
                    <Column style="padding-bottom: 10px">
                        <PasswordInput
                                bind:value={connection.token_secret}
                                labelText="Token Secret"
                                on:change={()=>connectionStore.persist(connection)}
                                required
                                type="password"
                        />
                    </Column>
                </Row>
            </Grid>
        </FluidForm>
    </Column>
</Row>

<script lang="ts">
    import {Column, FluidForm, Grid, PasswordInput, Row, TextInput} from "carbon-components-svelte";
    import {connectionStore} from "$lib/stores/connection";
    import {onMount} from "svelte";

    let connection = connectionStore.init();
    connectionStore.subscribe((c) => connection = c);
    onMount(async () => connection = await connectionStore.mount());

    /** handleLinkChange Set the Realm value based on the web services URL */
    const handleLinkChange = () => {
        connection.realm = "";
        if (!connection.ws_url) {
            return;
        }
        if (connection.ws_url.indexOf("//") === -1) {
            return;
        }
        const urlOne = connection.ws_url.split("//");
        const urlTwo = urlOne[1].split(".");
        if (!urlTwo[0]) {
            return;
        }
        connection.realm = urlTwo[0].toUpperCase().replace("-", "_");
    }

</script>