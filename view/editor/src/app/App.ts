import Vue from 'vue'
import Component from 'vue-class-component'
import BaseHeader from '@/components/BaseHeader/BaseHeader'
import AlertMessage from '@/components/AlertMessage/AlertMessage'
import navigation from '@/services/ServiceNavigation'
import session, {Session} from '@/services/ServiceSession'
import * as route from '@/router'

@Component({
	name: 'App',
	components: { BaseHeader, AlertMessage },
})
export default class App extends Vue {
	session:Session = session

	public mounted() {
		// TODO: implement ping method to start or not the application
		if (!session.user.isConnected()) {
			navigation.replaceView(route.landingPagePath)
		} else {
			navigation.replaceView(route.resumePagePath)
		}
	}
}
