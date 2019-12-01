import Vue from "vue"
import Component from "vue-class-component"

/**
 * Component to edit a text with html editor
 */
@Component({
	props: {
		text: { type: String, default: "", required: true },
		textChangedFunction: { type: Function, required: true },
	},
})
export default class TextEditor extends Vue {
	change(input: string) {
		this.$props.textChangedFunction(input)
	}
}
