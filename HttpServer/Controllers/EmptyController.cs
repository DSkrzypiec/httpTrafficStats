using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace HttpServer.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class EmptyController : Controller
    {
        [HttpGet]
        public void Get()
        {
        }
    }
}
