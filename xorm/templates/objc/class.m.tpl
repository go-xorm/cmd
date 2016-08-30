{{range .Tables}}
#import "{{UnTitle (Mapper .Name)}}.h"
#import "../NSDictionary+json.h"

@implementation {{Mapper .Name}}
{{$name := .Name}}
+({{Mapper .Name}} *)paserJson:(NSString *)jsonString{
    NSError *error = nil;
    if (jsonString == nil){
        return nil;
    }
    id jsonObject = [NSJSONSerialization JSONObjectWithData:[jsonString dataUsingEncoding:NSUTF8StringEncoding] options:NSJSONReadingMutableContainers error:&error];
    if (jsonObject != nil && [jsonObject isKindOfClass:[NSDictionary class]]) {
        NSDictionary *json = (NSDictionary*)jsonObject;

        {{Mapper .Name}} *{{.Name}} = [[[{{Mapper .Name}} alloc]init] autorelease];
        {{range .Columns}}{{if or (eq (Type .) "int") (eq (Type .) "long")}}{{if eq .Name "id"}}{{$name}}.remote_id = [[json valueForKey:@"{{Mapper .Name}}"] longValue];
        {{else}}{{$name}}.{{Mapper .Name}} = [[json valueForKey:@"{{Mapper .Name}}"] longValue];
        {{end}}{{else if eq (Type .) "BOOL"}}{{$name}}.{{Mapper .Name}} = [[json valueForKey:@"{{Mapper .Name}}"] boolValue];
        {{else if eq (Type .) "float"}}{{$name}}.{{Mapper .Name}} = [[json valueForKey:@"{{Mapper .Name}}"] floatValue];
        {{else if eq (Type .) "double"}}{{$name}}.{{Mapper .Name}} = [[json valueForKey:@"{{Mapper .Name}}"] doubleValue];
        {{else}}{{$name}}.{{Mapper .Name}} = [json valueForKey:@"{{Mapper .Name}}"];
        {{end}}{{end}}
        return {{$name}};
    }
    
    return nil;
}

-(NSString *)toJson{
    NSMutableDictionary *json = [[NSMutableDictionary alloc]init];
    
    {{range .Columns}}{{if or (eq (Type .) "int") (eq (Type .) "long") (eq (Type .) "float") (eq (Type .) "double")}}{{if eq .Name "id"}}[json setValue:[NSNumber numberWithLong:self.remote_id] forKey:@"{{Mapper .Name}}"];
    {{else}}[json setValue:[NSNumber numberWithLong:self.{{Mapper .Name}}] forKey:@"{{Mapper .Name}}"];
    {{end}}{{else}}[json setValue:self.{{Mapper .Name}} forKey:@"{{Mapper .Name}}"];
    {{end}}{{end}}
    
    NSError *error = nil;
    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:json options:kNilOptions error:&error];
    [json release];
    if (jsonData == nil){
        return @"";
    }
    
    return [[NSString alloc]initWithData:jsonData encoding:NSUTF8StringEncoding];
}

-(void)dealloc{
    
    {{range .Columns}}{{if eq (Type .) "NSString*"}}[_{{Mapper .Name}} release];
    {{end}}{{end}}
    [super dealloc];
}

@end

{{end}}