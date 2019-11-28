import Vue from 'vue';
import Component from 'vue-class-component';
import userService from '@/services/user.service';
import navService from '@/services/ServiceNavigation';
import httpService from '@/services/http.service';

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

  private getBook() {
    this.getBookLaunch = true;
    const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
    httpService.get(`api/books/${this.$route.params.id}`, headers, (response:any) => {
      this.book = response.data;
      this.nodes = [];
      this.book.nodeIds.forEach((nodeId:string) => {
        const node = {
          identifier: nodeId,
        };
        this.nodes.push(node);
        this.getNode(node);
      });
      this.getBookLaunch = false;
      console.log(JSON.stringify(this.book));
    }, (error:any) => {
      this.getBookLaunch = false;
    });
  }

  /**
   * Get node and replace in nodes list
   * @param node
   */
  public getNode(node:any) {
    const localNode = node;
    localNode.getLaunch = true;
    const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
    httpService.get(`api/books/${this.book.identifier}/node/${localNode.identifier}`, headers, (response:any) => {
      const newNode = response.data;
      const indexNode = this.nodes.findIndex(n => n.identifier === newNode.identifier);
      this.nodes.splice(indexNode, 1, newNode);
    }, (error:any) => {
      localNode.getLaunch = false;
    });
  }

  public createNode() {
    this.createNodeLaunch = true;
    const headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userService.getToken()}`);
    httpService.post(`api/books/${this.book.identifier}/node`, {
      title: 'New node',
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
