import { Component, Output, EventEmitter } from '@angular/core'
import { FileUploader, FileItem } from 'ng2-file-upload'
import { File } from '../../domain/file';

@Component({
  selector: 'hd-upload',
  templateUrl: 'hd-upload.html'
})
export class HdUploadComponent {
  @Output() up = new EventEmitter<File>()
  uploader: FileUploader = new FileUploader({ url: '/up' });
  constructor() {
    this.uploader.onCompleteItem = (item: FileItem, r: string, status: number) => {
      if (status !== 200) {
        // TOOD 服务器错误
        return
      }
      const ret: any = JSON.parse(r)
      if (ret['success']) {
        this.up.emit(ret['file'])
      }
      item['f'] = ret['file']
    }
  }
}
