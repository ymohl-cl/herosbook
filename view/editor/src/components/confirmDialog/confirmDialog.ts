import Vue from 'vue';
import Component from 'vue-class-component';

@Component({
  props: {
    open: { type: Boolean, default: false },
    title: { type: String, default: 'Confirmer?' },
    description: { type: String, default: '' },
    confirmCallback: Function,
    dismissCallback: Function,
  },
})
export default class ConfirmDialog extends Vue {
  public dismiss() {
    this.$props.dismissCallback();
  }

  public confirm() {
    this.$props.confirmCallback();
  }
}
