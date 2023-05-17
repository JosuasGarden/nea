<template>
  <div>
    <div class="items product-provider">
      <div class="block box" v-for="(item, n) in data.product_provider" :key="n">
        <div class="columns">
          <div class="column is-2">
            <b-field :label="$t('globals.buttons.enabled')">
              <b-switch v-model="item.enabled" name="enabled" :native-value="true" />
            </b-field>
            <b-field>
              <a
                @click.prevent="$utils.confirm(null, () => removeProvider(n))"
                href="#"
                class="is-size-7"
              >
                <b-icon icon="trash-can-outline" size="is-small" />
                {{ $t("globals.buttons.delete") }}
              </a>
            </b-field>
          </div>
          <!-- first column -->

          <div class="column" :class="{ disabled: !item.enabled }">
            <div class="columns">
              <div class="column is-4">
                <b-field :label="$t('globals.fields.type')" label-position="on-border">
                  <b-select v-model="item.type" name="type">
                    <option v-for="(_, type) in item.config" :key="type" :value="type">
                      {{ type }}
                    </option>
                  </b-select>
                </b-field>
                <b-field
                  v-for="(_, key) in item.config[item.type]"
                  :key="key"
                  :label="$t(`settings.productProvider.${item.type}.${key}`)"
                  label-position="on-border"
                  :message="$t(`settings.productProvider.${item.type}.${key}.Help`)"
                >
                  <b-input
                    v-model="item.config[item.type][key]"
                    :name="key"
                    :placeholder="placeholders[key]"
                    :maxlength="200"
                    required
                  />
                </b-field>
              </div>
            </div>
          </div>
        </div>
        <!-- second container column -->
      </div>
      <!-- block -->
    </div>
    <!-- product-provider -->

    <b-button @click="addProdiver" icon-left="plus" type="is-primary">
      {{ $t("globals.buttons.addNew") }}
    </b-button>
  </div>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

const placeholders = {
  endpoint: 'https://example.com',
  consumer_key: 'ck_xxxxxxxxxxxxxxxx',
  consumer_secret: 'cs_xxxxxxxxxxxxxxxx',
};

export default Vue.extend({
  props: {
    form: {
      type: Object,
    },
  },

  data() {
    return {
      data: this.form,
      placeholders,
    };
  },

  methods: {
    addProdiver() {
      this.data.product_provider.push({
        enabled: true,
        type: this.serverConfig.product_provider[0].type,
        config: Object.fromEntries(
          this.serverConfig.product_provider.map((provider) => [
            provider.type,
            JSON.parse(JSON.stringify(provider.config)),
          ]),
        ),
      });

      this.$nextTick(() => {
        const items = document.querySelectorAll('.product-provider select[name="type"]');
        items[items.length - 1].focus();
      });
    },

    removeProvider(i) {
      this.data.product_provider.splice(i, 1);
    },
  },

  computed: {
    ...mapState(['serverConfig', 'loading']),
  },
});
</script>
