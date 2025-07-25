# 项目需求文档
## 1. 项目背景
当前AI发展如火如荼，MCP（Model Context Protocol）作为一种新兴的协议，旨在为AI模型提供更高效的上下文管理和工具调用能力。我们的项目也在快速推进现有接口的MCP改造，但是在这个过程中，我们面临着一些挑战和问题，例如如何高效地将现有接口适配到MCP协议，如何保证改造后的接口性能和稳定性等。因此，产生了对这个问题进行梳理和解决的诉求。

## 2. 项目目标
本项目旨在构建一个MCP工具市场管理平台，主要目标包括：
1. **OPENAPI文档导入与转换**：支持导入符合OpenAPI规范的API文档，并将其解析转换为可通过MCP协议暴露的工具。
  1.1 **API文档解析**：能够解析OpenAPI规范的文档格式，提取出API的相关信息。
  1.2 **API测试**：提供API测试功能，确保转换后的API能够正常工作。
  1.3 **API管理**：支持对导入的API进行管理，包括查看、编辑和删除等操作。
2. **MCP服务集成**：内置MCP-Server，用于工具注册与管理。
  2.1 **MCP-Server生成**：根据用户需要自动生成MCP-Server。
  2.2 **工具注册**：支持将转换后的API工具注册到MCP-Server。
  2.3 **工具管理**：针对MCP-Server提供工具的管理功能，包括查看、编辑和删除等操作。
3. **MCP市场管理**：提供一个MCP市场管理平台，用户可以在平台上浏览、搜索和使用各种MCP工具。
  3.1 **工具浏览**：支持用户浏览市场中的所有MCP工具。
  3.2 **工具搜索**：提供搜索功能，用户可以根据关键词快速找到所需的MCP工具。
  3.3 **工具申请**：用户可以申请使用市场中的MCP工具，并进行相应的权限控制以及MCP的配置管理。

## 3. 技术方案
本项目将基于Golang 1.23开发，主要技术方案包括：
1. **Swagger文档解析**：使用开源库如`go-swagger`或`swaggo`来解析OpenAPI规范的Swagger文档。
    - `go-swagger`：提供了丰富的功能来处理Swagger文档，包括生成代码和文档等。
    - `swaggo`：轻量级的Swagger文档生成工具，支持注释方式生成文档。
2. **MCP协议适配**：根据MCP协议的要求，对转换后的API进行适配和封装。
    - 使用现有的MCP协议库，或根据需要自定义实现MCP协议的适配层。
    - 确保转换后的API能够通过MCP协议进行调用和管理。
3. **MCP-Server集成**：可以针对每个用户，提供一个独立的MCP-Server实例，支持多种暴露方式，如SSE（Server-Sent Events）、标准输入输出（stdio）等。
    - 使用Golang的网络编程能力，构建MCP-Server服务。
    - 提供RESTful API接口，供前端调用和管理MCP工具。
4. **前端展示**：使用现代前端框架（如React或Vue.js）来构建用户界面，实现API文档的导入、转换和MCP工具的管理功能。
    - 前端界面应简洁易用，支持API文档的上传、转换和MCP工具的浏览、搜索等功能。
    - 实现用户友好的交互体验，确保用户能够方便地使用MCP工具市场管理平台。
5. **数据库设计**：使用关系型数据库（如MySQL或PostgreSQL）存储API文档、MCP工具和用户信息等。
    - 设计合理的数据库表结构，确保数据的完整性和一致性。
    - 实现数据的增删改查操作，支持API文档和MCP工具的管理。

## 4. 执行计划
1. **需求分析**：详细梳理现有接口的MCP改造需求，明确各项功能的实现细节。
2. **技术选型**：确定使用的技术栈和工具，包括文档解析库、MCP协议适配方案等
3. **原型设计**：设计MCP工具市场管理平台的原型，包括前端界面和后端服务的交互流程。
4. **开发实现**：按照设计方案进行开发，实现各项功能模块。
5. **测试验证**：对开发完成的功能进行全面测试，确保其性能和稳定性。
6. **部署上线**：将项目部署到生产环境，并进行上线前的准备工作。
7. **维护与迭代**：上线后持续维护项目，收集用户反馈，进行功能迭代和优化。  


## 5. 预期成果
1. 成功构建一个MCP工具市场管理平台，支持OpenAPI文档导入与转换。
2. 实现MCP-Server的自动生成和工具注册功能。
3. 提供一个用户友好的界面，支持MCP工具的浏览、搜索和申请使用。
4. 确保平台的性能和稳定性，能够满足实际使用需求。


## 6. 参考资料
- [MCP协议文档](https://mcp-protocol.org/docs)
- [OpenAPI规范](https://swagger.io/specification/)
- [Golang官方文档](https://golang.org/doc/)
- [go-swagger](https://github.com/go-swagger/go-swagger)
- [swaggo](https://github.com/swaggo/swaggo)
- [React](https://reactjs.org/)
- [Vue.js](https://vuejs.org/)

## 7. 附录
- **项目时间表**：详细列出各阶段的时间节点和里程碑
- **团队成员分工**：明确各个团队成员的职责和任务分配   
- **风险评估与应对措施**：识别项目可能面临的风险，并制定相应的应对策略
- **技术文档**：提供项目相关的技术文档和API接口文档

## 8. 其他需求
- **用户权限管理**：实现用户权限的细粒度控制，确保不同用户对MCP工具的访问权限。
- **日志记录与监控**：集成日志记录和监控系统，便于问题排查和性能优化。
- **API文档生成**：自动生成API文档，方便用户查看和使用MCP工具。
- **多语言支持**：考虑到国际化需求，平台应支持多语言界面和文档。
- **安全性考虑**：确保平台的安全性，包括数据加密、身份验证等措施。