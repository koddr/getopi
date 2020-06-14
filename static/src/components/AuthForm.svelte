<script>
  import { goto } from "@sapper/app";
  import { getTokenData } from "../utils/jwt";

  // ...
  function handleSubmit(event) {
    // ...
    getTokenData(event.target.email.value, event.target.password.value);
  }
</script>

<style>

</style>

<form class="auth-form" on:submit|preventDefault={handleSubmit}>
  <input type="email" name="email" class="auth-form__input" required />
  <input type="password" name="password" class="auth-form__input" required />
  <button type="submit" class="auth-form__button">Submit</button>
</form>
