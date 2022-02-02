// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.OData.Deltas;
using Microsoft.AspNetCore.OData.Query;
using Microsoft.AspNetCore.OData.Routing.Controllers;
using Microsoft.EntityFrameworkCore;
using Microsoft.Identity.Web.Resource;
using ToDoApi.Models;

namespace ToDoApi.Controllers
{
    [Authorize]
    [RequiredScope("ToDoItem.ReadWrite")]
    public class ToDoItemsController : ODataController
    {
        private readonly ToDoContext _context;
        private readonly ILogger<ToDoItemsController> _logger;

        public ToDoItemsController(ToDoContext context, ILogger<ToDoItemsController> logger)
        {
            _context = context;
            _logger = logger;
        }

        // GET: /ToDoItems
        [HttpGet]
        [EnableQuery()]
        public async Task<ActionResult<IQueryable<ToDoItem>>> GetToDoItemsAsync()
        {
            await LogRequestAsync(Request);
            return _context.ToDoItems;
        }

        // GET: /ToDoItems/{id}
        [HttpGet]
        [EnableQuery()]
        public async Task<ActionResult<ToDoItem>> GetToDoItemAsync(Guid key)
        {
            await LogRequestAsync(Request);
            var todoItem = await _context.ToDoItems.FindAsync(key);

            if (todoItem == null)
            {
                return NotFound();
            }

            return todoItem;
        }

        // PUT: /ToDoItems/{id}
        [HttpPatch]
        public async Task<IActionResult> PatchToDoItemAsync(Guid key, Delta<ToDoItem> delta)
        {
            await LogRequestAsync(Request);
            if (delta.GetChangedPropertyNames().Contains("Id"))
            {
                return BadRequest();
            }

            var item = await _context.ToDoItems.FindAsync(key);

            if (item == null)
            {
                return NotFound();
            }

            delta.Patch(item);

            _context.Update(item);

            try
            {
                await _context.SaveChangesAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (!ToDoItemExists(key))
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return Updated(item);
        }

        // POST: /ToDoItems
        [HttpPost]
        public async Task<ActionResult<ToDoItem>> PostToDoItemAsync([FromBody]ToDoItem todoItem)
        {
            await LogRequestAsync(Request);
            _context.ToDoItems.Add(todoItem);
            await _context.SaveChangesAsync();

            return Created(todoItem);
        }

        // DELETE: /ToDoItems/{id}
        [HttpDelete]
        public async Task<IActionResult> DeleteToDoItemAsync(Guid key)
        {
            await LogRequestAsync(Request);
            var todoItem = await _context.ToDoItems.FindAsync(key);
            if (todoItem == null)
            {
                return NotFound();
            }

            _context.ToDoItems.Remove(todoItem);
            await _context.SaveChangesAsync();

            return NoContent();
        }

        private bool ToDoItemExists(Guid id)
        {
            return _context.ToDoItems.Any(e => e.Id == id);
        }

        private async Task LogRequestAsync(HttpRequest request)
        {
            var requestInfo = $"Incoming request: \n{request.Method} {request.Path}{request.QueryString}\n";

            if (request.ContentLength > 0 && request.HasJsonContentType())
            {
                request.Body.Position = 0;
                var reader = new StreamReader(request.Body);
                var payload = await reader.ReadToEndAsync();
                requestInfo += $"\n{payload}";
            }

            _logger.LogInformation(requestInfo);
        }
    }
}
