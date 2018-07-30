import { Component } from '@angular/core'
import { FileUploader } from 'ng2-file-upload'

const URL = '/up'
@Component({
  selector: 'hd-upload',
  templateUrl: 'hd-upload.html'
})
export class HdUploadComponent {
  public uploader: FileUploader = new FileUploader({ url: URL });
  constructor() {
  }
}
