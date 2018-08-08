import { Pipe, PipeTransform } from '@angular/core';

const types = new Map<String, String>();
types.set('image', '照片')
types.set('video', '视频')
types.set('audio', '音频')
types.set('archive', '归档文件')
@Pipe({
  name: 'fileType',
})
export class FileTypePipe implements PipeTransform {
  transform(value: string, ...args) {
    if (types.has(value)) {
      return types.get(value)
    }
    return '未知'
  }
}
