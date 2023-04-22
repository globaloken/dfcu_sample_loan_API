<template>
  <v-main>
    <v-row align="center" justify="center" style="height: 100vh">
      <v-card min-width="500px" elevation="0">
        <v-card class="pa-5">
          <v-card-title class="mb-5"><v-spacer /> Sign In <v-spacer /></v-card-title>
          <!-- <v-divider /> -->
          <v-card-text class="pa-0 ma-0">
            <v-form v-model="isValid" ref="loginForm" lazy-validation>
              <v-text-field v-model="form.username" label="Username" filled />
              <v-text-field v-model="form.password" label="Password" filled type="password" elevation="0" />
            </v-form>
          </v-card-text>
          <!-- <v-divider /> -->
          <v-card-actions>
            <v-spacer />
            <v-btn block color="primary" class="text-caption" dark @click="login" :loading="loading" large
              elevation="0">Login</v-btn>
          </v-card-actions>
        </v-card>
      </v-card>
    </v-row>

    <v-snackbar v-model="snackBar.model" :color="snackBar.color" right bottom>
      {{ snackBar.message }}
    </v-snackbar>

  </v-main>
</template>

<script>
import axios from "axios";

export default {
  name: 'Login',

  data: () => ({
    form: {
      username: "",
      password: "",
    },
    loading: false,
    isValid: false,
    snackBar: {
      model: false,
      color: "red",
      message: ""
    }
  }),

  methods: {
    async login() {
      this.loading = true;

      try {
        const response = await axios.post('/users/login', this.form);
        localStorage.setItem(process.env.VUE_APP_ACCESS_TOKEN_KEY, response.data.access_token);

        this.$router.replace("/home");
      } catch (e) {
        this.snackBar.message = "Invalid Username or Password"
        this.snackBar.model = true
      } finally {
        this.loading = false;
      }
    }
  }
}
</script>
