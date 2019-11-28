import httpService from '@/services/http.service';

export default class Book {
	identifier: string = "";
	title: string;
	description: string;
	genre: string;
	publish: boolean = false;
	ownerID: string = "";
	nodesID: string[] = [];
	creationDate: Date = new Date(0);
	categoriesID: string[] = [];

	constructor(title: string, description: string, genre: string) {
		this.title = title;
		this.description = description;
		this.genre = genre;
	}
	unmarshall(json: any) {
		this.identifier = json.identifier
		this.publish = json.publish
		this.ownerID = json.ownerID
		this.nodesID = json.nodesID
		this.creationDate = json.creationDate
	}
	record(userToken: string, callbackSucess: () => void) {
		let headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userToken}`);
		httpService.post('api/books', this, headers, (resp: any) => {
			this.unmarshall(resp.data)
			callbackSucess()
		}, (error:any) => {
			console.log("error")
			console.log(error)
		});
	}
	delete(userToken:string, callbackSucess: () => void) {
		let headers = httpService.appendHeaders(httpService.getDefaultHeaders(), 'Authorization', `Bearer ${userToken}`);
		httpService.delete(`api/books/${this.identifier}`, headers, (response:any) => {
			callbackSucess()
		}, (error:any) => {
			console.log("error")
			console.log(error)
		});
	}
}
