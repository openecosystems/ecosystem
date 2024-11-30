import warnings
from pprint import pprint
import argparse
import  urllib.parse, os
from urllib import parse;
#import SOAPpy
#from suds.client import Client
import sys
import io;
import os;


'''
Prefixes:
Global elements:
Global types:
Binding:
Service: name
   Port: portName
      Operations: opNmae(name: namespace:type,....) -> return: ns9:TaskInfo
'''





inBuiltConversion = {
  "boolean":"bool",
  "Any":"google.protobuf.Any",
  "dateTime":"google.protobuf.Timestamp",
  "long":"int64",
  "int":"int32",
  "base64Binary":"bytes"
}

class GrpcPackage:
  def __init__(self,url):
    spl =  url.split(":",1)
    self.ns = url.split(":")[0]
    self.packageName = CovertFromURL(url) if len(spl)>1 and spl[1].strip().startswith("http") else url;
    self.knownMessage = {}
    self.packageimports = set([])
    self.services = {}
  def ifExist(self,typename):
    return (typename in self.knownMessage)

  def AddMessage(self,typename,message,namespace):
    if not self.ifExist(typename):
      self.knownMessage[typename] = message
      self.packageimports.add(namespace)

  def AddService(self,name,grpcService):
    if not self.ifExist(name):
      self.services[name] = grpcService
  def GetMessageText(self):
    text=""
    for name,msg  in self.knownMessage.items():
      text+= (msg.GetText(1)+"\n");
    return text;

  def GetServiceText(self,opname):
    text = ""
    for name,srv  in self.services.items():
      text+= (srv.GetText(opname)+"\n");
    return text;

  def GetText(self):
    text = "\nsyntax = \"proto3\";\n\n"
    text +=  ("package {};\n\n".format(self.packageName));
    for imp in self.packageimports:
      if imp is not None and imp.strip() not in "xsd":
        if packages[imp] is not self:
          text+= ("import \"{}\";".format(packages[imp].packageName+".proto"))
          text+="\n"
    text += self.GetMessageText();
    text += self.GetServiceText(1);
    return text;





def CovertFromURL(url):
  if url.endswith("xsd"):
    comps = url.split("/")[2];
    comps = comps.split('.');
    comps.reverse()
    namespace = ".".join(comps);
    url = "http://"+namespace;
  urlParsed = urllib.parse.urlparse(url)
  return "{0}{1}".format(urlParsed.netloc,''.join(filter(str.isalnum ,urlParsed.path)))

def GetPackage(name):
  global packages;
  if name in packages:
    return packages[name]
  pack = GrpcPackage(name);
  packages[name] =pack;
  return pack

def uncapitalize(s):
  return s[:1].lower() + s[1:]

class GrpcField:
  def __init__(self,con):
    self.packName = False;
    self.isBuiltIn = False;
    if len(con.split(":")) > 2 :
      spl = con.split(":",1);
      self.varname = uncapitalize(spl[0])
      typeSpl = spl[1].split(":",1)
      self.package = typeSpl[0].strip()
      if spl[1].endswith("[]"):
        self.typename = "Any" if "None" in typeSpl[1] else  typeSpl[1].replace("[]","")
        self.repeated = True;
      else:
        self.typename = "Any" if "None" in typeSpl[1] else  typeSpl[1]
        self.repeated = False;
    elif con.startswith("xsd") or con.startswith("ns"):
      self.package = con.split(":",1)[0].strip()
      spl = con.split(":",1);
      self.varname = uncapitalize(spl[0])
      if spl[1].endswith("[]"):
        self.typename = "Any" if "None" in spl[1] else  spl[1].replace("[]","")
        self.repeated = True;
      else:
        self.typename = "Any" if "None" in spl[1] else  spl[1]
        self.repeated = False;
    else:
      spl =  con.split(":",1)
      self.package = None
      spl = con.split(":",1);
      self.varname = spl[0]
      if spl[1].endswith("[]"):
        self.typename = "Any" if "None" in spl[1] else  spl[1].replace("[]","")
        self.repeated = True;
      else:
        self.typename = "Any" if "None" in spl[1] else  spl[1]
        self.repeated = False;
    if self.typename in inBuiltConversion:
      self.isBuiltIn= True;
      self.typename = inBuiltConversion[self.typename]
    if self.package is not None and 'xsd' in self.package:
      self.isBuiltIn= True;

  def GetText(self,indent):
    if self.repeated:
      return "\t repeated {} {}".format(self.typename  if (not self.packName or self.isBuiltIn or self.package is None) else packages[self.package].packageName+"."+self.typename,self.varname);
    else:
      return "\t {} {}".format(self.typename  if (not self.packName or self.isBuiltIn or self.package is None) else packages[self.package].packageName+"."+self.typename,self.varname);

#TopLevel Request will have name Op+"Request" if Not Null otherwise None
class GrpcMessage:
  def __init__(self,msg,isWrap = False,service=None):
    if not isWrap:
      split =  msg.split(":",1)
      global packages
      self.package = packages[split[0]]
      self.message = split[1]
      if self.message.endswith(")"):
        fieldPart = self.message[self.message.find("(")+1:self.message.find(")")]
        self.name = self.message.replace("("+fieldPart+")","")
        self.acceptableMessage = []
        self.isEnum = False;
        self.fields = [];
        for f in fieldPart.split(","):
          if f != '':
            self.AddField(f);
      else:
        self.name = split[1]
        self.isEnum = True;
      self.package.AddMessage(self.name,self,None);
    else:
      self.isEnum = False;
      self.fields = [];
      self.package = service.GetPackage();

  def AddField(self,fdstring):
    f = GrpcField(fdstring.strip())
    if "google.protobuf.Timestamp" in f.typename:
      self.package.packageimports.add("time");
    elif "google.protobuf.Empty" in f.typename:
      self.package.packageimports.add("empty");
    elif "google.protobuf.Any" in f.typename:
      self.package.packageimports.add("any");
    else:
      self.package.packageimports.add(f.package)
    self.fields.append(f);

  def SetAcceptableMessage(self,message):
    self.acceptableMessage.append(message)

  def GetText(self,indent):
    if not self.isEnum:
      text = "message {} {{\n".format(self.name);
      i=1;
      for f in self.fields:
        if not f.package ==  self.package.ns:
          f.packName =  True

        text += (f.GetText(1)+"={};\n".format(i))
        i+=1
      return text+"}"
    else:
      text = "enum {} {{\n".format(self.name);
      return text+"}"

class GrpcOperation:
  def __init__(self,opname,service):
    opsplt = opname.split("->",1);
    if len(opsplt) > 1:
      self.output = opsplt[1];
    else:
      self.output = None
    params =  opsplt[0][opsplt[0].find("(")+1:opsplt[0].find(")")]
    self.name =  opsplt[0].replace(params,"").replace("()","").strip();
    self.params = params.split(",")
    self.params = [i for i in self.params if i]
    self.service = service
    if len(self.params)>0:
      self.req = GrpcMessage(None,isWrap=True,service=self.service)
      self.req.name = self.name+"Request";
      self.service.GetPackage().AddMessage(self.req.name,self.req,None);
      for p in self.params:
        self.req.AddField(p);
    else:
      self.params = None
      self.service.GetPackage().packageimports.add("empty")
    if self.output is not None:
      self.output = self.output.strip()
    if  self.output is not None and self.output not in "None":
      self.res = GrpcMessage(None,isWrap=True,service=self.service)
      self.res.name = self.name+"Response";
      self.outparams = self.output.split(",")
      self.outparams = [i for i in self.outparams if i]
      for outp in self.outparams:
        self.res.AddField(outp);
      self.service.GetPackage().AddMessage(self.res.name,self.res,None);
    else:
      self.output = None;
      self.service.GetPackage().packageimports.add("empty")
  def GetText(self,indent):
    return ('\t'* indent) + "rpc {}({}) returns ({}) {{}}".format(self.name ,self.name + "Request" if self.params is not None else 'google.protobuf.Empty', self.name+ "Response" if self.output is not None else 'google.protobuf.Empty')

class GrpcService:
  def __init__(self,name):
    self.servicename = name
    self.operations = {}
    self.servMessage = {}
    self.package = GrpcPackage(name);
    packages[name] = self.package;
    self.package.AddService(name,self)
  def AddOperation(self,operation,overide=False):
    if operation.name in self.operations and not overide:
      #warnings.warn('Duplicate operation '+ operation.name,UserWarning)
      return;
    self.operations[operation.name] = operation;
  def GetPackage(self):
    return packages[self.servicename]
  def GetText(self,operation):
    text = "service {} {{\n".format(self.servicename)
    for opn, op in self.operations.items():
      text += (op.GetText(1)+"\n");
    return text +"}"

def BuildGrpc():
  for  name,pack in packages.items():
    import os
    if not os.path.exists("out"):
      os.makedirs("out")

    if name != "xsd" and not pack.packageName.startswith("google"):
      with open(os.path.join("./out",pack.packageName)+".proto", "w") as f:
        f.write(pack.GetText())




def parse_indentation(raw_python_file_contents):
  """Parse the contents of a file with python style indentation.
  >>> clean_file_contents = '''
  ... # ignored line
  ... level 1 # comments are ignored
  ...     level 2
  ...         level 3
  ... level 1
  ... level 1
  ...     level 2'''
  >>> out = parse_python_indentation(clean_file_contents)
  >>> result = [{
  ...     'key': 'level 1',
  ...     'offspring': [{
  ...         'key': 'level 2',
  ...         'offspring': [{
  ...             'key': 'level 3',
  ...             'offspring': []
  ...         }]
  ...     }]
  ... }, {
  ...     'key': 'level 1',
  ...     'offspring': []
  ... }, {
  ...     'key': 'level 1',
  ...     'offspring': [{
  ...         'key': 'level 2',
  ...         'offspring': []
  ...     }]
  ... }]
  >>> cmp(out, result)
  0
  >>> unclean_file_contents = '''
  ... # ignored line
  ... level 1 # comments are ignored
  ...     level 2
  ...       level 3 # The indentation level is too low here.
  ... level 1
  ... level 1
  ...     level 2'''
  >>> with warnings.catch_warnings(record=True) as w:
  ...     warnings.simplefilter("always")
  ...     a = parse_python_indentation(unclean_file_contents)
  >>> len(w)
  1
  >>> w[0].category
  <type 'exceptions.UserWarning'>
  >>> str(w[0].message)
  'Indentation with errors!'
  """

  raw_lines = raw_python_file_contents.split('\n')
  cleaned_lines = []
  python_output = []
  indentation_length = 0
  error = False
  i = 0

  # Remove comments and empty lines
  for raw_line in raw_lines:
    line_wo_comment = raw_line.split('#')[0].rstrip()
    if len(line_wo_comment.strip()) > 0:
      cleaned_lines.append(line_wo_comment)

  # Find indentation length
  while indentation_length == 0 and i < len(cleaned_lines):
    if len(cleaned_lines[i]) - len(cleaned_lines[i].lstrip()) > 0:
      indentation_length = len(cleaned_lines[i]) - len(cleaned_lines[i].lstrip())
    i += 1

  # Don't allow indentations of zero
  if indentation_length == 0:
    indentation_length = 1

  # Turn Python into a construct of Lists and Objects
  for cleaned_line in cleaned_lines:
    indentations = float(len(cleaned_line) - len(cleaned_line.lstrip())) / indentation_length
    current_list = python_output

    if indentations % 1 != 0:
      # indentation characters do not correspond to a known indentation level.
      error = True

    indentations = int(indentations)

    for j in range(0, indentations):
      if len(current_list) == 0:
        # The indentations tell us to go a place that it's not possible to go.
        error = True
      else:
        current_list = current_list[-1]['offspring']

    current_list.append({
      'key': cleaned_line.lstrip(),
      'offspring': []
    })
  return python_output




builtIn = ["string","long","base64Binary","int","dateTime","boolean"]
serviceMap = {}

defaultXSD = GrpcPackage("xsd")
googleAny = GrpcPackage("google/protobuf/any")
googleEmpty = GrpcPackage("google/protobuf/empty")
googleTime = GrpcPackage("google/protobuf/timestamp")
endpoint = None
file = None;
packages = {"any":googleAny,
            "empty":googleEmpty,
            "time":googleTime}
if __name__ == "__main__":
  try:
    from zeep import Client
  except ImportError:
    print('Error: zeep is not installed! use pip install zeep')
    exit(1);
  if len(sys.argv) < 3:
    print('Usage: python3 convert.py wsdl binding')
    exit(1)
  else:
    file = sys.argv[1]
    endpoint  = sys.argv[2]
  import subprocess
  alltext = subprocess.getoutput("python -mzeep {}".format(file))
  zeepTree = parse_indentation(alltext);
  #resolve packages
  allpackage = zeepTree[0];
  for items in allpackage['offspring']:
    if not items['key'].startswith("xsd"):
      pack= GrpcPackage(items['key'])
      packages[pack.ns] = pack;
    else:
      packages['xsd'] =defaultXSD;

  allGlobalElements = zeepTree[1];
  for items in allGlobalElements['offspring']:
    if not items['key'].startswith("xsd"):
      msg = GrpcMessage(items['key'])

  allGlobalType = zeepTree[2];
  for items in allGlobalType['offspring']:
    if not items['key'].startswith("xsd"):
      msg = GrpcMessage(items['key'])
  allServices = zeepTree[4];
  serviceName = allServices["key"].split(":",1)[1].strip().capitalize()
  serv = GrpcService(serviceName);
  serviceMap[serviceName] = serv
  serviceFound = False
  outFile = False;
  for items in allServices['offspring']:
    if endpoint in items['key']:
      serviceFound = True;
      continue;
    if serviceFound is True:
      outFile = True;
      for op in items['offspring']:
        grpcOp = GrpcOperation(op['key'],serv)
        serv.AddOperation(grpcOp)
      serviceFound = False
  #for msg in serv.GetPackage().knownMessage.values():
  #print( serv.package.GetText());
  if outFile:
    BuildGrpc()
  else:
    print("Could not find binding {}".format(endpoint));
    exit(1)
