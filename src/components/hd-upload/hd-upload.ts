import { Component, Output, EventEmitter } from '@angular/core'
import { FileUploader } from 'ng2-file-upload'
import { File } from '../../domain/file';

@Component({
  selector: 'hd-upload',
  templateUrl: 'hd-upload.html'
})
export class HdUploadComponent {
  @Output() up = new EventEmitter<File>()
  uploader: FileUploader = new FileUploader({ url: '/up' });
  constructor() {
    this.uploader.response.subscribe((r: string) => {
      if (r.length < 10) { return }
      const f: File = JSON.parse(r)
      for (const q of this.uploader.queue) {
        if (q.file && q.file.name == f.name) {
          q.alias = f.id
        }
      }
      this.up.emit(f)
    })
  }
}
