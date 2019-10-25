import Vue from 'vue';
import Component from 'vue-class-component';
import textEditor from '../../components/textEditor/textEditor.vue';

import userService from '../../../services/user.service';
import navService from '../../../services/nav.service';
import httpService from '../../../services/http.service';

@Component
export default class Book extends Vue {

	public book:any = false;
	public nodes:any[] = [];
	public getBookLaunch:boolean = false;
	public createNodeLaunch:boolean = false;

	public mounted() {
		if (!userService.isConnected() || this.$route.params.id == null) {
			navService.replaceView('/login');
		} else {
			this.getBook();
		}
	}

	private getBook(){
		this.getBookLaunch = true;
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
		console.log(headers);
		httpService.get('api/books/' + this.$route.params.id, headers, (response:any) => {
			this.book = response.data;
			this.nodes = [];
			this.book.nodeIds.map((nodeId:string) => {
				var node = {
					identifier:nodeId
				}
				this.nodes.push(node);
				this.getNode(node);
			});
			this.getBookLaunch = false;
			console.log(JSON.stringify(this.book));
		}, (error:any) => {
			this.getBookLaunch = false;
		});
	}

	public getNode(node:any){
		node.getLaunch = true;
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
		httpService.get('api/books/' + this.book.identifier + '/node/' + node.identifier, headers, (response:any) => {
			console.log(JSON.stringify(response.data));
			var newNode = response.data;
			var indexNode = this.nodes.findIndex(node => node.identifier == newNode.identifier);
			this.nodes[indexNode] = newNode;
		}, (error:any) => {
			node.getLaunch = false;
		});
	}

	public createNode(){
		this.createNodeLaunch = true;
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
		httpService.post('api/books/' + this.book.identifier + '/node', {
			title:'New node',
		}, headers, (response:any) => {
			console.log(JSON.stringify(response.data));
			this.nodes.push(response.data);
			this.book.nodeIds.push(response.data.identifier);
			this.createNodeLaunch = false;
			console.log(JSON.stringify(this.book));
		}, (error:any) => {
			this.createNodeLaunch = false;
		});
	}
	
}
