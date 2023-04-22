<template>
  <v-main>
    <v-row class="pa-5">
      <v-col cols="12" sm="3">
        <v-card>
          <v-card-title>Requests</v-card-title>
          <v-card-text>
            <span>{{ this.requests }}</span>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="3">
        <v-card>
          <v-card-title>Failed validations</v-card-title>
          <v-card-text>
            <span>{{ this.failed_validations }}</span>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="3">
        <v-card>
          <v-card-title>Positive Requests</v-card-title>
          <v-card-text>
            <span>{{ this.postive_requests }}</span>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="3">
        <v-card>
          <v-card-title>Negative Requests</v-card-title>
          <v-card-text>
            <span>{{ this.negative_requests }}</span>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-main>
</template>

<script>
import axios from "axios";

export default {
  name: 'Home',

  data: () => ({
    loading: false,
    requests: 0,
    failed_validations: 0,
    postive_requests: 0,
    negative_requests: 0,
    snackBar: {
      model: false,
      color: "red",
      message: ""
    }
    // number of requests, number of failed validations, number of positive requests (at least one loan) and number of negative requests 
  }),

  methods: {
    async loadLogs() {
      this.loading = true;

      try {
        const token = localStorage.getItem(process.env.VUE_APP_ACCESS_TOKEN_KEY);

        const response = (await axios.get('/logs', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })).data;

        const positive = response.find((x) => x.type === 'POSITIVE_REQUEST')?.count || 0;
        const negative = response.find((x) => x.type === 'NEGATIVE_REQUEST')?.count || 0;
        const request = response.find((x) => x.type === 'REQUEST')?.count || 0;
        const failed_validation = response.find((x) => x.type === 'FAILED_VALIDATION')?.count || 0;

        this.requests = request;
        this.failed_validations = failed_validation;
        this.postive_requests = positive;
        this.negative_requests = negative;

      } catch (e) {
        console.log(e);
        this.snackBar.message = "Error loading logs"
        this.snackBar.model = true
      } finally {
        this.loading = false;
      }
    }
  },

  mounted() {
    this.loadLogs();
  }


}
</script>
