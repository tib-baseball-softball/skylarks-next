<script lang="ts">
  import {Carta, MarkdownEditor} from 'carta-md';
  import 'carta-md/default.css';
  import DOMPurify from 'dompurify';

  interface Props {
    value: string;
    required?: boolean;
  }

  let {value = $bindable(), required = false}: Props = $props();

  const carta = new Carta({
    sanitizer: DOMPurify.sanitize,
  });
</script>

<MarkdownEditor bind:value mode="tabs" {carta} textarea={{required: required}}/>
