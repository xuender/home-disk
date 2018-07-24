import { Component } from '@angular/core';

/**
 * Generated class for the HdUploadComponent component.
 *
 * See https://angular.io/api/core/Component for more info on Angular
 * Components.
 */
@Component({
  selector: 'hd-upload',
  templateUrl: 'hd-upload.html'
})
export class HdUploadComponent {

  text: string;

  constructor() {
    console.log('Hello HdUploadComponent Component');
    this.text = 'Hello World';
  }

}
