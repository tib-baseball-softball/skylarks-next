<script lang="ts">
  import {Mail, User} from "lucide-svelte";
  import Avatar from "$lib/dp/components/utils/Avatar.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    model: CustomAuthModel;
  }

  const {model}: Props = $props();
</script>

<div class="card preset-tonal-surface user-card shadow-lg">
  <header class="card-header">
    <h2 class="h4 card-title">User Data</h2>
  </header>

  <section class="user-content">
    <Avatar
      --size="4rem"
      fallback={`${model.first_name.charAt(0)?.toUpperCase()}${model.last_name.charAt(0)?.toUpperCase()}`}
      src={client.files.getURL(model, model?.avatar)}
    />

    <div class="user-details">
      <div class="info-row">
        <User/>
        <div>
          <p>{`${model?.first_name ?? ""} ${model?.last_name ?? ""}`}</p>
          <p class="info-label">Name</p>
        </div>
      </div>

      <div class="info-row">
        <Mail/>
        <div>
          <p>{model?.email}</p>
          <p class="info-label">email address</p>
        </div>
      </div>
    </div>
  </section>

  <footer class="card-footer">
    <p>Edit User Data:</p>
    <a class="btn preset-tonal-primary border border-primary-500" href="/account/{model.id}/settings"
       title="Go to User Settings">User Settings</a>
  </footer>
</div>

<style>
  .user-card {
    @media (min-width: 64rem) {
      grid-column: span 2 / span 2;
    }
  }

  .card-title {
    font-weight: var(--font-weight-semibold);
  }

  .user-content {
    padding: calc(var(--spacing) * 4);
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 4);

    @media (min-width: 40rem) {
      flex-direction: row;
      align-items: center;
      gap: calc(var(--spacing) * 12);
    }
  }

  .user-details {
    display: grid;
    grid-template-columns: repeat(1, minmax(0, 1fr));
    gap: calc(var(--spacing) * 2);
    grid-column: span 4 / span 4;
  }

  .info-row {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 3);
  }

  .info-label {
    font-size: var(--text-sm);
    line-height: var(--tw-leading, var(--text-sm--line-height));
    font-weight: var(--font-weight-light);
  }

  .card-footer {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    gap: calc(var(--spacing) * 4);
  }
</style>
