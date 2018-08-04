import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
@Injectable()
export class FilesProvider {
  constructor(public http: HttpClient) {
    console.log('Hello FilesProvider Provider');
  }
}
