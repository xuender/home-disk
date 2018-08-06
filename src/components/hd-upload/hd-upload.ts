import { Component, Output, EventEmitter } from '@angular/core'
import { FileUploader, FileItem } from 'ng2-file-upload'

@Component({
  selector: 'hd-upload',
  templateUrl: 'hd-upload.html'
})
export class HdUploadComponent {
  @Output() up = new EventEmitter<boolean>()
  uploader: FileUploader = new FileUploader({ url: '/up' });
  constructor() {
  }
  uploadAll() {
    this.uploader.uploadAll()
    this.up.emit(true)
  }
  upload(item: FileItem) {
    item.upload()
    this.up.emit(true)
  }
}
