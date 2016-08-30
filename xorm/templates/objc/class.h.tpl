#import <Foundation/Foundation.h>

{{range .Tables}}@interface {{Mapper .Name}} : NSObject

{{range .Columns}}@property(nonatomic,{{if eq (Type .) "NSString*"}}retain{{else}}assign{{end}}) {{Type .}} {{if eq .Name "id"}}remote_id{{else}}{{Mapper .Name}}{{end}};
{{end}}
+({{Mapper .Name}} *)paserJson:(NSString *)jsonString;
-(NSString *)toJson;

@end
{{end}}